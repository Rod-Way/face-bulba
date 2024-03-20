import React, { useState } from 'react';
import axios from 'axios';

const Uploader = () => {
	const [selectedFile, setSelectedFile] = useState(null);

	const handleFileChange = event => {
		setSelectedFile(event.target.files[0]);
	};

	const handleUpload = () => {
		if (!selectedFile) {
			alert('Please select a file!');
			return;
		}

		const formData = new FormData();
		formData.append('file', selectedFile);

		axios
			.post('http://localhost:5000/api/catch-data', formData, {
				headers: {
					'Content-Type': 'multipart/form-data',
				},
			})
			.then(response => {
				// Handle success
				console.log('Photo uploaded successfully:', response.data);
				alert('Photo uploaded successfully!');
			})
			.catch(error => {
				// Handle error
				console.error('Error uploading photo:', error);
				alert('Error uploading photo. Please try again.');
			});
	};

	return (
		<div>
			<input type='file' onChange={handleFileChange} />
			<button onClick={handleUpload}>Upload</button>
		</div>
	);
};

export default Uploader;
