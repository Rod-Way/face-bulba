import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import Footer from './components/Footer/Footer';
import NavBar from './components/NavBar/NavBar';
import Login from './pages/LoginPage';
import Main from './pages/Main';
import Register from './pages/RegisterPage';
import './styles/main.css';
import ScrollToTop from './utils/scrollToTop';

function App() {
	return (
		<div className='App'>
			<Router>
				<ScrollToTop />

				<NavBar />
				<Routes>
					<Route path='/' element={<Main />} />
					<Route path='/login' element={<Login />} />
					<Route path='/register' element={<Register />} />
				</Routes>
				<Footer />
			</Router>
		</div>
	);
}

export default App;
