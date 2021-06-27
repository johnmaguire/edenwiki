import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import ReactMarkdown from 'react-markdown';

import ErrorMessage from '../components/ErrorMessage';

export default function Page() {
  const [isErrored, setIsErrored] = useState(false);
  const [isLoaded, setIsLoaded] = useState(false);
  const [page, setPage] = useState<{Body: string}>({Body: ""});
  const { pageName } = useParams<{pageName: string}>();

  useEffect(() => {
    fetch("http://localhost:3000/page/"+pageName)
      .then(res => res.json())
      .then(
        (result) => {
          setIsLoaded(true);
          setPage(result);
        },
        (error) => {
          console.error(error);
          setIsLoaded(true);
          setIsErrored(true);
        }
      )
  }, [pageName]);

  let children: JSX.Element = <></>;
  if (!isLoaded) {
    children = <p>Loading page...</p>;
  } else if(isErrored) {
    children = <ErrorMessage>Unable to load the page.</ErrorMessage>;
  } else {
    children = <ReactMarkdown>{page.Body}</ReactMarkdown>
  }

  return (
    <>
      <h2>{pageName}</h2>

      {children}
    </>
  )
}
