import React, { useState } from 'react';
import { NavLink, Link } from 'react-router-dom';
import { setCookie } from '../utils/useCookie';

const Login = () => {
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');

	const [jwtToken, setJwtToken] = useState('');

	const handleSubmit = event => {
		fetch('http://localhost:5000/api/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({
				user: username,
				password: password,
			}),
		})
			.then(response => {
				if (!response.ok) {
					throw new Error('Failed to login');
				}
				return response.json();
			})
			.then(data => {
				console.log(data);
				if (data.token) {
					setCookie('jwt', data.token, {
						secure: true,
						httpOnly: true,
					});
				} else {
					throw new Error('Token not found');
				}
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
					autoComplete='current-username'
					onChange={e => setUsername(e.target.value)}
					required
				/>
				<input
					type='password'
					placeholder='Пароль'
					value={password}
					autoComplete='current-password'
					onChange={e => setPassword(e.target.value)}
					required
				/>
				<button type='submit'>
					{/* <Link to='/'>Отправить</Link> */}
				</button>
				<NavLink className='other' to='/register'>
					Нет аккаунта? <strong>Зарегистрироваться</strong>
				</NavLink>
			</form>
		</div>
	);
};

export default Login;
