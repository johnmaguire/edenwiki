import { NavLink } from 'react-router-dom';

import styles from './BottomNav.module.css';

export default function BottomNav() {
  return (
    <nav>
      <ul className={styles.ul}>
        <li className={styles.li}><NavLink to="/new" activeClassName={styles.active}>Create a new page</NavLink></li>
      </ul>
    </nav>
  );
}
