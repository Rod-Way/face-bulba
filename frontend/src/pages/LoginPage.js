import React, { useState } from 'react';

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
		<div id='login'>
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
			</form>
			<button>
				Нет аккаунта? <strong>Зарегистрироваться</strong>
			</button>
		</div>
	);
};

export default Login;
