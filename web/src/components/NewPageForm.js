import { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { useForm } from 'react-hook-form';

export default function NewPageForm() {
  const [pageName, setPageName] = useState("");
  const [submitError, setSubmitError] = useState("");
  const { register, handleSubmit, formState: { errors } } = useForm();

  const onSubmit = (data) => {
    fetch("http://localhost:3000/page/" + data.title, {method: "PUT", body: JSON.stringify({Body: data.body})})
      .then(
        () => {
          setPageName(data.title);
        },
        (error) => {
          setSubmitError(error.message);
        },
      )
  }

  return pageName ? <Redirect to={"/page/"+pageName} /> : (
    <>
      {submitError && <p>Error: {submitError}</p>}
      <form onSubmit={handleSubmit(onSubmit)}>
        <p>
          <label>Title: <input {...register("title", { required: true })} /></label>
          {errors.title && <span>A title is required.</span>}
        </p>
        <div>
          <textarea {...register("body", { required: true })}></textarea>
          {errors.body && <span>A title is required.</span>}
        </div>
        <input type="submit" name="submit" value="Create" />
      </form>
    </>
  );
}
