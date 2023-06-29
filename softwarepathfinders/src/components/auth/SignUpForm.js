import React, { useState, useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux'
import { Navigate } from 'react-router-dom';
import { signUp } from '../../store/session';

const SignUpForm = () => {
  const [errors, setErrors] = useState([]);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [repeatPassword, setRepeatPassword] = useState('');
  const [first_name, setFirst_Name] = useState('');
  const [last_name, setLast_Name] = useState('');
  const [phone, setPhone] = useState('');
  const user = useSelector(state => state.session.user);
  const dispatch = useDispatch();
  const [submitted, setSubmitted] = useState(false);

  const emailRegX = /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/

  useEffect(() => {
    let errors = [];

    if (email.length < 2 || email.length > 50 || email.includes(" ")) {
      errors.push("Email must be between 2 and 50 characters and contain no white space")
    }
    if (!email.match(emailRegX)) {
      errors.push("Email must be a valid email address ( example@ex.com )")
    }
    if (password.length < 6 || password.length > 30 || password.includes(" ")) {
      errors.push('Password must be between 6 and 30 characters and contain no white space');
    }
    if (first_name.length < 1 || first_name.length > 20 || first_name.includes(" ")) {
      errors.push("First name must be between 1 and 20 characters and contain no white space")
    }
    if (last_name.length < 1 || last_name.length > 20 || last_name.includes(" ")) {
      errors.push("Last name must be between 1 and 20 characters and contain no white space")
    }
    if (password !== repeatPassword) {
      errors.push('Passwords must match');
    }

    setErrors(errors);
  }, [email, password, repeatPassword, first_name, last_name, phone]);



  const onSignUp = async (e) => {
    e.preventDefault();
    setSubmitted(true);

    if (errors.length > 0) return
    if (password === repeatPassword) {
      const data = await dispatch(signUp(email, password, first_name, last_name, phone));
      if (data) {
        setErrors(data)
      }
    }
  };



  const updateEmail = (e) => {
    setEmail(e.target.value);
  };

  const updatePassword = (e) => {
    setPassword(e.target.value);
  };

  const updateFirstname = (e) => {
    setFirst_Name(e.target.value);
  };

  const updateLastname = (e) => {
    setLast_Name(e.target.value);
  };
  const updateRepeatPassword = (e) => {
    setRepeatPassword(e.target.value);
  };

  const updatePhone = (e) => {
    setPhone(e.target.value);
  };

  if (user) {
    return <Navigate to='/login' />;
  }

  return (
    <div className="form-outer-container">
      <form onSubmit={onSignUp}>
      <div className="form-header">Please Fill Out The Following Fields:</div>
      <div className="required-field">(Fields labeled with&nbsp;<div className="asterisk">*</div>&nbsp;are required)</div>
      <div className="required-field">(Fields must not contain any white space)</div>
        <div className="form-container">
          <div className="create_errors">
            {submitted && (errors).map((error, i) => (
              <div className="errorMessageContainer" key={i}>
                <div className="errorMessage">{error}</div>
              </div>
            ))}
          </div>
          <div>
            <label className='form-field-label'>Email&nbsp;<div className="asterisk">*</div></label>
            <input
              className="form-field"
              type='text'
              name='email'
              placeholder='Email'
              onChange={updateEmail}
              value={email}
              required
            ></input>
          </div>
          <div>
            <label className='form-field-label'>Password&nbsp;<div className="asterisk">*</div></label>
            <input
              className="form-field"
              type='password'
              name='password'
              placeholder='Password'
              onChange={updatePassword}
              value={password}
              required
            ></input>
          </div>
          <div>
            <label className='form-field-label'>Confirm Password&nbsp;<div className="asterisk">*</div></label>
            <input
              className="form-field"
              type='password'
              name='repeat_password'
              placeholder='Confirm Password'
              onChange={updateRepeatPassword}
              value={repeatPassword}
              required
            ></input>
          </div>
          <div>
            <label className='form-field-label'>First Name&nbsp;<div className="asterisk">*</div></label>
            <input
              className="form-field"
              type='text'
              name='first_name'
              placeholder='First Name'
              onChange={updateFirstname}
              value={first_name}
              required
            ></input>
          </div>
          <div>
            <label className='form-field-label'>Last Name&nbsp;<div className="asterisk">*</div></label>
            <input
              className="form-field"
              type='text'
              name='last_name'
              placeholder='Last Name'
              onChange={updateLastname}
              value={last_name}
              required
            ></input>
          </div>
          <div>
            <label className='form-field-label'>Phone&nbsp;<div className="asterisk">*</div></label>
            <input
              className="form-field"
              type='phone'
              name='phone'
              placeholder='Phone'
              onChange={updatePhone}
              value={phone}
              required
            ></input>
          </div>
          <button type='submit' className="form-button">Sign Up</button>
        </div>
      </form>
    </div>
  );
};

export default SignUpForm;

// auth branch change
