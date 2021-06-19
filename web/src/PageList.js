import { useEffect, useState } from 'react';

function PageList() {
  const [error, setError] = useState(null);
  const [isLoaded, setIsLoaded] = useState(false);
  const [pages, setPages] = useState([]);

  useEffect(() => {
    fetch("http://localhost:3000/page")
      .then(res => res.json())
      .then(
        (result) => {
          setIsLoaded(true);
          console.log(result);
          setPages(result);
        },
        (error) => {
          setIsLoaded(true);
          setError(error);
        }
      )
  }, []);

  if (!isLoaded) {
    return <p>Loading pages...</p>
  } else if(error !== null) {
    return <p>Error: {error.message}</p>
  } else {
    if (Object.keys(pages).length === 0) {
      return <p>No pages exist yet. Create one?</p>
    }

    return (
      <ul>
        {Object.keys(pages).map((page, i) => (
          <li>
            <a href="#" key={page}>{page}</a>
          </li>
        ))}
      </ul>
    );
  }
}

export default PageList;
