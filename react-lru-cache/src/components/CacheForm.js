import React, { useState } from 'react';
import { setCache } from '../utils/Myapp';

const CacheForm = () => {
    const [key, setKey] = useState('');
    const [value, setValue] = useState('');
    const [expiration, setExpiration] = useState(5);

    const handleSubmit = async (e) => {
        e.preventDefault();
        await setCache({ key, value, expiration });
        setKey('');
        setValue('');
        setExpiration(5);
    };

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label>Key: </label>
                <input type="text" value={key} onChange={(e) => setKey(e.target.value)} />
            </div>
            <div>
                <label>Value: </label>
                <input type="text" value={value} onChange={(e) => setValue(e.target.value)} />
            </div>
            <div>
                <label>Expiration (seconds): </label>
                <input type="number" value={expiration} onChange={(e) => setExpiration(e.target.value)} />
            </div>
            <button type="submit">Set Cache</button>
        </form>
    );
};

export default CacheForm;
