import React, { useState } from 'react';

const Register = () => {
	const [formData, setFormData] = useState([
		{ label: 'Имя', type: 'text', value: '' },
		{ label: 'Фамилия', type: 'text', value: '' },
		{ label: 'Логин', type: 'text', value: '' },
		{ label: 'Электронная почта', type: 'email', value: '' },
		{ label: 'Пароль', type: 'password', value: '' },
	]);

	const handleSubmit = event => {
		event.preventDefault();
		// TODO: SEND DATA TO BACKEND
		formData.forEach(field =>
			console.log(`${field.label}: ${field.value}`)
		);
	};

	const handleChange = (index, value) => {
		const newFormData = [...formData];
		newFormData[index].value = value;
		setFormData(newFormData);
	};

	return (
		<div id='register'>
			<form id='register-form' onSubmit={handleSubmit}>
				{formData.map((field, index) => (
					<input
						key={index}
						type={field.type}
						placeholder={field.label}
						value={field.value}
						onChange={e => handleChange(index, e.target.value)}
					/>
				))}
				<button id='register-btn' type='submit'>
					Отправить
				</button>
				<button id='sign-up'>
					Уже есть аккаунт? <strong>Войти</strong>
				</button>
			</form>
		</div>
	);
};

export default Register;
