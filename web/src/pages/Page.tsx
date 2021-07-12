import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import ReactMarkdown from 'react-markdown';
import gfm from 'remark-gfm';
// @ts-ignore
import wikiLink from 'remark-wiki-link';

import ErrorMessage from '../components/ErrorMessage';


export default function Page() {
  const [isErrored, setIsErrored] = useState<boolean>(false);
  const [isLoaded, setIsLoaded] = useState<boolean>(false);
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

  return (
    <>
      <h2>{pageName}</h2>

      {isLoaded ? (
        isErrored ? <ErrorMessage>Unable to load the page.</ErrorMessage> :
          <ReactMarkdown remarkPlugins={
            [[gfm],
             [wikiLink, {
               aliasDivider: '|',
               pageResolver: (name: string) => [name],
               hrefTemplate: (permalink: string) => `/page/${permalink}`,
             }]]}>{page.Body}</ReactMarkdown>
       ) : (
        <p>Loading page...</p>
      )}
    </>
  )
}
