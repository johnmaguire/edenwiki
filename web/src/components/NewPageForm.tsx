import { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { FormProvider, useForm } from 'react-hook-form';

import ErrorMessage from '../components/ErrorMessage';
import MarkdownPreviewTextarea from '../components/MarkdownPreviewTextarea'
import styles from './NewPageForm.module.css';

export default function NewPageForm() {
  const [pageName, setPageName] = useState<string>("");
  const [isErrored, setIsErrored] = useState<boolean>(false);
  const methods = useForm();
  const { register, handleSubmit, formState: { errors } } = methods;

  const onSubmit = (data: {title: string, body: string}) => {
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
    <FormProvider {...methods}>
      {isErrored && <ErrorMessage>Failed to create the page.</ErrorMessage>}
      <form onSubmit={handleSubmit(onSubmit)}>
        <p>
          <label>Title: <input className={styles.title} {...register("title", { required: true })} /></label>
          {errors.title && <ErrorMessage>A title is required.</ErrorMessage>}
        </p>

        <div>
          <MarkdownPreviewTextarea name="body" />
        </div>

        <input type="submit" name="submit" value="Create" />
      </form>
    </FormProvider>
  );
}
