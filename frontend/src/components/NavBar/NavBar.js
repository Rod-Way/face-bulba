import { NavLink } from 'react-router-dom';
import './style.css';
import BtnDarkMode from '../btnDarkMode/BtnDarkMode';

const NavBar = () => {
	return (
		<div id='navbar'>
			<BtnDarkMode />
			<h1 className='title-2'>Account:</h1>
			<p>
				Username: <br />
			</p>
			<NavLink to='/'>К основной странице</NavLink>
			<NavLink to='/userS'>Искать юзеров</NavLink>
			<NavLink to='/user/:userID'>Страница</NavLink>
		</div>
	);
};

export default NavBar;
