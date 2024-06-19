package cache

import (
    "sync"
    "time"
    "go-lru-cache/internals/list"
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
    items    map[string]*list.Node
    order    *list.DoublyLinkedList
    lock     sync.RWMutex
}

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        items:    make(map[string]*list.Node),
        order:    list.NewDoublyLinkedList(),
    }
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
    c.lock.RLock()
    defer c.lock.RUnlock()

    if node, found := c.items[key]; found {
        item := node.Value.(*CacheItem)
        if item.isExpired() {
            return nil, false
        }
        c.order.MoveToFront(node)
        return item.value, true
    }
    return nil, false
}

func (c *LRUCache) Set(key string, value interface{}, duration time.Duration) {
    c.lock.Lock()
    defer c.lock.Unlock()

    if node, found := c.items[key]; found {
        c.order.MoveToFront(node)
        node.Value.(*CacheItem).value = value
        node.Value.(*CacheItem).expiration = time.Now().Add(duration).UnixNano()
        return
    }

    if len(c.items) >= c.capacity {
        c.evict()
    }

    item := &CacheItem{
        key:        key,
        value:      value,
        expiration: time.Now().Add(duration).UnixNano(),
    }

    node := &list.Node{
        Key:   key,
        Value: item,
    }

    c.order.PushFront(node)
    c.items[key] = node
}

func (c *LRUCache) Delete(key string) {
    c.lock.Lock()
    defer c.lock.Unlock()

    if node, found := c.items[key]; found {
        c.order.remove(node)
        delete(c.items, key)
    }
}

func (c *LRUCache) evict() {
    node := c.order.RemoveTail()
    if node != nil {
        delete(c.items, node.Key)
    }
}

func (c *LRUCache) GetAllItems() map[string]interface{} {
    c.lock.RLock()
    defer c.lock.RUnlock()

    items := make(map[string]interface{})
    current := c.order.head
    for current != nil {
        item := current.Value.(*CacheItem)
        if !item.isExpired() {
            items[current.Key] = item.value
        }
        current = current.Next
    }
    return items
}

