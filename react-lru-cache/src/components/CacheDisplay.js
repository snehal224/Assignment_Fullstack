
import React, { useState, useEffect } from 'react';
import { getCache, deleteCache } from '../utils/Myapp';

const CacheDisplay = () => {
    const [key, setKey] = useState('');
    const [value, setValue] = useState(null);

    const handleGet = async () => {
        const data = await getCache(key);
        setValue(data.value);
    };

    const handleDelete = async () => {
        await deleteCache(key);
        setValue(null);
        setKey('');
    };

    return (
        <div>
            <div>
                <label>Key: </label>
                <input type="text" value={key} onChange={(e) => setKey(e.target.value)} />
                <button onClick={handleGet}>Get Cache</button>
                <button onClick={handleDelete}>Delete Cache</button>
            </div>
            {value && (
                <div>
                    <h3>Value: {value}</h3>
                </div>
            )}
        </div>
    );
};

export default CacheDisplay;

