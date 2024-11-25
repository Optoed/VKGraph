document.getElementById('findPath').addEventListener('click', async () => {
    const user1 = document.getElementById('user1').value.trim();
    const user2 = document.getElementById('user2').value.trim();
    const resultDiv = document.getElementById('result');

    if (!user1 || !user2) {
        resultDiv.textContent = 'Пожалуйста, введите оба ID!';
        return;
    }

    resultDiv.textContent = 'Поиск пути...';

    try {
        const response = await fetch(`/friends/info/${user1}/${user2}`);
        if (!response.ok) {
            throw new Error(`Ошибка: ${response.statusText} (${response.status})`);
        }

        const data = await response.json();

        if (!Array.isArray(data) || data.length === 0) {
            resultDiv.textContent = 'Путь не найден или данные некорректны.';
            return;
        }

        resultDiv.innerHTML = `
            <h2>Кратчайший путь:</h2>
            <ul>
                ${data.map(user => `
                    <li>
                        <img src="${user.photo}" alt="${user.name}" style="width: 50px; height: 50px; border-radius: 50%; margin-right: 10px;">
                        <a href="${user.source}" target="_blank">${user.name}</a> (ID: ${user.id})
                    </li>
                `).join('')}
            </ul>
        `;
    } catch (error) {
        console.error('Ошибка при выполнении запроса:', error);
        resultDiv.textContent = `Ошибка: ${error.message}`;
    }
});
