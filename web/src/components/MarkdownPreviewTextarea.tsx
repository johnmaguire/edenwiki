import React from 'react';
import ReactMarkdown from 'react-markdown';
import gfm from 'remark-gfm';
import { Control, useFormContext, useWatch } from 'react-hook-form';

import ErrorMessage from '../components/ErrorMessage';
import styles from './MarkdownPreviewTextarea.module.css';

type Props = React.DetailedHTMLProps<React.TextareaHTMLAttributes<HTMLTextAreaElement>, HTMLTextAreaElement> & {
  className?: string;
  name: string;
}

function PreviewRenderer({ inputName, control }: {inputName: string, control: Control}) {
  const text = useWatch({
    control,
    name: inputName,
    defaultValue: "",
  });
  return <ReactMarkdown className={styles.preview_pane} remarkPlugins={[[gfm]]}>{text}</ReactMarkdown>;
}

export default function MarkdownPreviewTextarea(props: Props) {
  const { register, formState: { errors }, control } = useFormContext();

  return (
    <div className={styles.markdown_preview}>
      <div className={styles.textarea_wrapper}>
        <textarea className={styles.textarea} {...register(props.name, { required: true })} {...props}></textarea>
        {errors.body && <ErrorMessage>A body is required.</ErrorMessage>}
      </div>
      <PreviewRenderer inputName={props.name} control={control} />
    </div>
  );
};
