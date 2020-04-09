import React from 'react';
import axios from 'axios';

const App = () => {
  const getProfile = async () => {
    try {
      const result = await axios.get('/api/test', {
        headers: { 'Content-Type': 'application/json' },
        withCredentials: true,
      });
      console.log(result);
    } catch (error) {
      console.log('error!!');
    }
  };

  // const getProfile = async () => {
  //   fetch('http://localhost:8081/api/test', {
  //     mode: 'no-cors',
  //   })
  //     .then((res) => {
  //       res.text();
  //     })
  //     .then((text) => {
  //       console.log(text);
  //     });
  // };

  return (
    <div className="App">
      <div>
        <button onClick={() => getProfile()}>get profile!</button>
      </div>
    </div>
  );
};

export default App;
