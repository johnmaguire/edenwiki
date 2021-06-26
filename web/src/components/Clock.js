import { useState } from 'react';

function Clock() {
  const [time, setTime] = useState(new Date());
  setInterval(() => { setTime(new Date()) }, 1000);

  return (
    <p>
      The time is currently {time.toLocaleTimeString()}.
    </p>
  )
}

export default Clock;
