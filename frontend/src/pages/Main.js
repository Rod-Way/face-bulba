import Header from '../components/Header/Header';
import React, { useState, useEffect } from 'react';

const Main = () => {
	const [data, setData] = useState([]);
	const [loading, setLoading] = useState(false);
	const [prevData, setPrevData] = useState([]);
	const [batchNum, setBatchNum] = useState(1); // Сохраняем номер пакета в состоянии

	useEffect(() => {
		fetchMoreData();
	}, []); // Зависимость пуста, поэтому useEffect вызывается только после монтирования компонента

	const fetchMoreData = () => {
		setLoading(true);
		fetch(`http://localhost:5000/api/get/posts-batch/${batchNum}`)
			.then(response => response.json())
			.then(data => {
				const posts = data.response || [];
				if (JSON.stringify(posts) === JSON.stringify(prevData)) {
					alert('Новых данных нет (');
				} else {
					setData(prevData => [...prevData, ...posts]);
					setPrevData(posts);
				}
				setLoading(false);
				return;
			})
			.catch(error => {
				console.error('Error fetching data:', error);
				setLoading(false);
			});
	};

	const handleButtonClick = () => {
		setBatchNum(prevBatchNum => prevBatchNum + 1);
		fetchMoreData();
	};

	return (
		<div id='main'>
			<Header />
			<div className='container'>
				{data.map((post, index) => (
					<div className='card' key={index}>
						{console.log(post)}
						<p className='author'>Author: {post.author}</p>
						<p className='text'>Text: {post.text}</p>
						<p className='files'>Files: {post.files_url}</p>
						<p>
							Тэги:{' '}
							{post.tags.map((tag, index) => (
								<p key={index}>{tag}</p>
							))}
						</p>
						<p>
							Комментарии:{' '}
							{post.comments.map((comment, index) => (
								<div className='card comment' key={index}>
									<p>Author: {comment.author}</p>
									<p>Text: {comment.text}</p>
									<p>Создано {post.createdat}</p>
								</div>
							))}
						</p>
						{post.is_updated && <p>Данные обновлены</p>}
						<p>Создано {post.createdAt}</p>
					</div>
				))}
				{loading && <div>Loading...</div>}
				<button className='card-button' onClick={handleButtonClick}>
					Получить больше
				</button>
			</div>
		</div>
	);
};

export default Main;
