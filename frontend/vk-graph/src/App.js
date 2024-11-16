import React, { useEffect, useRef } from 'react';
import { Network } from 'vis-network';
import 'vis-network/styles/vis-network.css';

const App = ({ users }) => {
  const networkRef = useRef(null);

  useEffect(() => {
    // Проверка на наличие данных
    if (!users || users.length === 0) return;

    const nodes = users.map(user => ({
      id: user.id,
      label: user.name,
      shape: 'circularImage',
      image: user.photo,
      title: `<a href="${user.source}" target="_blank">${user.name}</a>`,
    }));

    const edges = users.slice(1).map((user, index) => ({
      from: users[index].id,
      to: user.id,
    }));

    const container = networkRef.current;
    const data = { nodes, edges };
    const options = {
      nodes: {
        borderWidth: 2,
        size: 50,
        color: {
          border: '#2B7CE9',
          background: '#97C2FC',
        },
        font: { color: '#343434' },
      },
      edges: {
        color: '#848484',
        arrows: { to: { enabled: true, scaleFactor: 1.2 } },
      },
      interaction: { hover: true },
    };

    new Network(container, data, options);
  }, [users]);

  // Проверка на случай, если данные ещё не загрузились
  if (!users || users.length === 0) {
    return <div>Загрузка...</div>;
  }

  return <div ref={networkRef} style={{ height: '600px', width: '100%' }} />;
};

export default App;
