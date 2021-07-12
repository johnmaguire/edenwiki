import {
  BrowserRouter as Router,
  Switch,
  Redirect,
  Route,
  Link,
} from 'react-router-dom';

import PageList from './pages/PageList';
import Page from './pages/Page';
import NewPage from './pages/NewPage';
import BottomNav from './components/BottomNav';

import './styles/global';

function App() {
  return (
    <Router>
      <div className="App">
        <h1><Link to="/">EdenWiki</Link></h1>

        <Switch>
          <Route exact path="/">
            <Redirect to="/page/Home" />
          </Route>
          <Route exact path="/page">
            <PageList />
          </Route>
          <Route path="/page/:pageName">
            <Page />
          </Route>
          <Route exact path="/new">
            <NewPage />
          </Route>
        </Switch>

        <BottomNav />
      </div>
    </Router>
  );
}

export default App;
