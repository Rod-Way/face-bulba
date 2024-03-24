import './style.css';
import gitHub from './gitHub.svg';
import telegram from './telegram.svg';

const Footer = () => {
	return (
		<footer className='footer'>
			<div className='container'>
				<div className='footer__wrapper'>
					<ul className='social'>
						<li className='social__item'>
							<a
								href='https://github.com/Rod-Way/face-bulba'
								target='_blank'
								rel='noopener noreferrer'
							>
								<img src={gitHub} alt='Link' />
							</a>
						</li>
						<li className='social__item'>
							<a
								href=''
								target='_blank'
								rel='noopener noreferrer'
							>
								<img src={telegram} alt='Link' />
							</a>
						</li>
					</ul>
					<div className='copyright'>
						<p>Это подвал</p>
					</div>
				</div>
			</div>
		</footer>
	);
};

export default Footer;
