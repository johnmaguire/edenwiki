import styles from './ErrorMessage.module.css';

type Props = {
    children: React.ReactNode;
}

export default function ErrorMessage({ children }: Props) {
  return (
    <div className={styles.ErrorMessage}>
      {children}
    </div>
  )
}
