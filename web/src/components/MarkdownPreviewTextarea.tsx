import React from 'react';
import { Control, useFormContext, useWatch } from 'react-hook-form';
import ReactMarkdown from 'react-markdown';
import gfm from 'remark-gfm';
// @ts-ignore
import wikiLink from 'remark-wiki-link';

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
  return (
    <ReactMarkdown className={styles.preview_pane} remarkPlugins={
      [[gfm],
       [wikiLink, {
         aliasDivider: '|',
         pageResolver: (name: string) => [name],
         hrefTemplate: (permalink: string) => `/page/${permalink}`,
       }]]}>{text}</ReactMarkdown>
  );
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
