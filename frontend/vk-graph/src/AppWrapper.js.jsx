import React, { useEffect, useState } from 'react';
import App from './App';

const AppWrapper = () => {
    const [users, setUsers] = useState([]);

    const fetchData = async () => {
        try {
            const response = await fetch('http://localhost:8080/friends/info/265240894/6482392');
            const data = await response.json();
            setUsers(data);
        } catch (error) {
            console.error('Ошибка при получении данных:', error);
        }
    };

    useEffect(() => {
        fetchData();
    }, []);

    return (
        <div>
            <h1>Визуализация графа друзей VK</h1>
            <App users={users} />
        </div>
    );
};

export default AppWrapper;
