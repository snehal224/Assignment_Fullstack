package cache

import (
    "container/list"
    "sync"
    "time"
)

type CacheItem struct {
    key        string
    value      interface{}
    expiration int64
}

func (item *CacheItem) isExpired() bool {
    return time.Now().UnixNano() > item.expiration
}

type LRUCache struct {
    capacity int
    items    map[string]*list.Element
    order    *list.List
    lock     sync.RWMutex
}

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        items:    make(map[string]*list.Element),
        order:    list.New(),
    }
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
    c.lock.RLock()
    defer c.lock.RUnlock()

    if element, found := c.items[key]; found {
        item := element.Value.(*CacheItem)
        if item.isExpired() {
            return nil, false
        }
        c.order.MoveToFront(element)
        return item.value, true
    }
    return nil, false
}

func (c *LRUCache) Set(key string, value interface{}, duration time.Duration) {
    c.lock.Lock()
    defer c.lock.Unlock()

    if element, found := c.items[key]; found {
        c.order.MoveToFront(element)
        element.Value.(*CacheItem).value = value
        element.Value.(*CacheItem).expiration = time.Now().Add(duration).UnixNano()
        return
    }

    if c.order.Len() >= c.capacity {
        c.evict()
    }

    item := &CacheItem{
        key:        key,
        value:      value,
        expiration: time.Now().Add(duration).UnixNano(),
    }

    element := c.order.PushFront(item)
    c.items[key] = element
}

func (c *LRUCache) Delete(key string) {
    c.lock.Lock()
    defer c.lock.Unlock()

    if element, found := c.items[key]; found {
        c.order.Remove(element)
        delete(c.items, key)
    }
}

func (c *LRUCache) evict() {
    element := c.order.Back()
    if element != nil {
        item := element.Value.(*CacheItem)
        delete(c.items, item.key)
        c.order.Remove(element)
    }
}
