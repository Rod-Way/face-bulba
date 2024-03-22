import React, { useState } from 'react';

import { NavLink } from 'react-router-dom';

const Login = () => {
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');

	const handleSubmit = event => {
		event.preventDefault();
		// TODO: SEND DATA TO BACKEND
		console.log('Username:', username);
		console.log('Password:', password);
	};
	return (
		<div className='card'>
			<form onSubmit={handleSubmit}>
				<input
					type='text'
					placeholder='Логин'
					value={username}
					onChange={e => setUsername(e.target.value)}
				/>
				<input
					type='password'
					placeholder='Пароль'
					value={password}
					onChange={e => setPassword(e.target.value)}
				/>
				<button type='submit'>Отправить</button>
				<NavLink className='other' to='/register'>
					Нет аккаунта? <strong>Зарегистрироваться</strong>
				</NavLink>
			</form>
		</div>
	);
};

export default Login;
