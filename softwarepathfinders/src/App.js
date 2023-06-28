// import React from 'react';
// import './App.css';

// function App() {
//   return (
//     <></>
//   )
// }

// export default App;
import React, { useState, useEffect } from 'react';
// import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { Provider } from 'react-redux';
import { useDispatch } from 'react-redux';
import SignUpForm from './components/auth/SignUpForm';
import { authenticate } from './store/session';
import './App.css';
import {store} from './store/store';

function App() {
  const [loaded, setLoaded] = useState(false);
  const dispatch = useDispatch();

  useEffect(() => {
    (async () => {
      await dispatch(authenticate());
      setLoaded(true);
    })();
  }, [dispatch]);

  // if (!loaded) {
  //   return null;
  // }

  return loaded && (
    <Router>
      <Provider store={store}>
        <Routes>
          <Route path='/sign-up' element={SignUpForm} exact={true}>
            <SignUpForm />
          </Route>
        </Routes>
      </Provider>
    </Router>
  );
}

export default App;
