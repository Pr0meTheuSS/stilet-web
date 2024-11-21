import React from 'react';
import DroneControlPage from './Pages/DroneControlPage';
import './App.css';
import ReactDOM from 'react-dom';
import { DroneServiceProvider } from './Services/DroneService';

const App = () => (
  <DroneServiceProvider>
      <DroneControlPage />
  </DroneServiceProvider>
  );


export default App;

ReactDOM.render(<App />, document.getElementById('root'));
