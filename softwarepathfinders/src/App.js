import React, { useState, useEffect } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
// import { Provider } from 'react-redux';
import { useDispatch } from 'react-redux';
import SignUpForm from './components/auth/SignUpForm';
import { authenticate } from './store/session';
import './App.css';
// import {store} from './store/store';
import NavBar from './components/navbar/NavBar';

function App() {
  const [loaded, setLoaded] = useState(false);
  const dispatch = useDispatch();

  useEffect(() => {
      setLoaded(true);
  }, [dispatch]);


  return loaded && (
    <BrowserRouter>
      {/* <Provider store={store}> */}
        <Routes>
          <Route path='/signup' element={<SignUpForm />} exact={true} />
        </Routes>
        <Routes>
          <Route path='/' element={<NavBar />} />

        </Routes>
      {/* </Provider> */}
    </BrowserRouter>
  );
}

export default App;
