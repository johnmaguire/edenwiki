import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
} from 'react-router-dom';

import Home from './pages/Home';
import Page from './pages/Page';
import NewPage from './pages/NewPage';

import './styles/global';

function App() {
  return (
    <Router>
      <div className="App">
        <h1><Link to="/">GardenWiki</Link></h1>

        <Switch>
          <Route exact path="/">
            <Home />
          </Route>
          <Route path="/page/:pageName">
            <Page />
          </Route>
          <Route path="/new">
            <NewPage />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
