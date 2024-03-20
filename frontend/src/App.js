import Footer from './components/Footer/Footer';
import NavBar from './components/NavBar/NavBar';
import Uploader from './components/upload/Uploader';
import Login from './pages/LoginPage';
import Main from './pages/Main';
import Register from './pages/RegisterPage';
import './styles/main.css';

function App() {
	return (
		<div className='App'>
			<NavBar />
			<Main />
			<Uploader />
			<Login />
			<Register />
			<Footer />
		</div>
	);
}

export default App;
