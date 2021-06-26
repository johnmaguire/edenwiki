import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

export default function Page() {
  const [error, setError] = useState(null);
  const [isLoaded, setIsLoaded] = useState(false);
  const [page, setPage] = useState({});
  const { pageName } = useParams();

  useEffect(() => {
    fetch("http://localhost:3000/page/"+pageName)
      .then(res => res.json())
      .then(
        (result) => {
          setIsLoaded(true);
          setPage(result);
        },
        (error) => {
          setIsLoaded(true);
          setError(error);
        }
      )
  }, [pageName]);

  let content = null;
  if (!isLoaded) {
    content = <p>Loading page...</p>;
  } else if(error !== null) {
    content = <p>Error: {error.message}</p>;
  } else {
    content = <p>{page.Body}</p>
  }

  return (
    <>
      <h2>{pageName}</h2>
      <p>{content}</p>
    </>
  )
}
