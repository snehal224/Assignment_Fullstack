export const setCache = async ({ key, value, expiration }) => {
    await fetch('http://localhost:8080/set', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ key, value, expiration }),
    });
};

export const getCache = async (key) => {
    const response = await fetch(`http://localhost:8080/get?key=${key}`);
    return await response.json();
};

export const deleteCache = async (key) => {
    await fetch(`http://localhost:8080/delete?key=${key}`, {
        method: 'DELETE',
    });
};
