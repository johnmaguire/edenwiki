import Clock from '../components/Clock';
import PageList from '../components/PageList';
import { Link } from 'react-router-dom';

export default function Page(props) {
  return (
    <>
      <h2>Home</h2>

      <Clock />

      <PageList />

      <p><Link to="/new">Create a new page</Link></p>
    </>
  )
}
