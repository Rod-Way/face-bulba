import Footer from './components/Footer/Footer';
import NavBar from './components/NavBar/NavBar';
import Uploader from './components/upload/Uploader';
import Login from './pages/LoginPage';

function App() {
	return (
		<div className='App'>
			<NavBar />
			<Uploader />
			<Login />
			<Footer />
		</div>
	);
}

export default App;
