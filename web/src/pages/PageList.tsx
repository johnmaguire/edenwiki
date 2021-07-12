import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

import ErrorMessage from '../components/ErrorMessage';

function PageList() {
  const [isErrored, setIsErrored] = useState<boolean>(false);
  const [isLoaded, setIsLoaded] = useState<boolean>(false);
  const [pages, setPages] = useState<string[]>([]);

  useEffect(() => {
    fetch("http://localhost:3000/page")
      .then(res => res.json())
      .then(
        (result) => {
          setIsLoaded(true);
          setPages(result.Pages);
        },
        (error) => {
          console.error(error);
          setIsLoaded(true);
          setIsErrored(true);
        }
      )
  }, []);

  if (!isLoaded) {
    return <p>Loading pages...</p>
  } else if(isErrored) {
    return <ErrorMessage>Unable to fetch available pages.</ErrorMessage>
  } else {
    if (pages.length === 0) {
      return <p>No pages exist yet. <Link to={"/new"}>Create one?</Link></p>
    }

    return (
      <>
        <h2>Pages</h2>
        <ul>
          {pages.map((page) => (
            <li>
              <Link to={"/page/"+page} key={page}>{page}</Link>
            </li>
          ))}
        </ul>
      </>
    );
  }
}

export default PageList;
