import React from 'react';
import CacheForm from './components/CacheForm';
import CacheDisplay from './components/CacheDisplay';

const App = () => {
    return (
        <div>
            <h1>LRU Cache</h1>
            <CacheForm />
            <CacheDisplay />
        </div>
    );
};

export default App;

