import React, { useState } from 'react';

import { NavLink } from 'react-router-dom';

const Login = () => {
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');

	const handleSubmit = event => {
		event.preventDefault();
		fetch(`http://localhost:5000/api/login`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Authorization: 'Bearer ваш_токен_авторизации', // Замените ваш_токен_авторизации на ваш реальный токен
			},
		})
			.then(response => response.json())
			.then(data => {
				return;
			})
			.catch(error => {
				console.error('Error fetching data:', error);
			});
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
