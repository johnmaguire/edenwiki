import { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { useForm } from 'react-hook-form';

import ErrorMessage from '../components/ErrorMessage';
import styles from './NewPageForm.module.css';

export default function NewPageForm() {
  const [pageName, setPageName] = useState("");
  const [isErrored, setIsErrored] = useState(false);
  const { register, handleSubmit, formState: { errors } } = useForm();

  const onSubmit = (data) => {
    fetch("http://localhost:3000/page/" + data.title, {method: "PUT", body: JSON.stringify({Body: data.body})})
      .then(
        () => {
          setPageName(data.title);
        },
        (error) => {
          console.error(error);
          setIsErrored(true);
        },
      )
  }

  return pageName ? <Redirect to={"/page/"+pageName} /> : (
    <>
      {isErrored && <ErrorMessage>Failed to create the page.</ErrorMessage>}
      <form onSubmit={handleSubmit(onSubmit)}>
        <p>
          <label>Title: <input className={styles.title} {...register("title", { required: true })} /></label>
          {errors.title && <ErrorMessage>A title is required.</ErrorMessage>}
        </p>
        <div>
          <textarea className={styles.body} {...register("body", { required: true })}></textarea>
          {errors.body && <ErrorMessage>A body is required.</ErrorMessage>}
        </div>
        <input type="submit" name="submit" value="Create" />
      </form>
    </>
  );
}
