import './App.css';
import { useState } from 'react';
import Container from 'react-bootstrap/Container';
import Card from 'react-bootstrap/Card';
import Navbar from 'react-bootstrap/Navbar';
import Alert from 'react-bootstrap/Alert';
import Result from './Result';
import RiddleForm from './RiddleForm';

const SUCCESS = 'success';
const NOT_FOUND = 'not_found';
const ERROR = 'error';

function App() {

  const [result, setResult] = useState({});
  const [status, setStatus] = useState('');
  const [error, setError] = useState('');

  const setSuccess = (value) => {
    setStatus(SUCCESS);
    setResult(value);
  }

  const setStatusError = (value) => {
    setStatus(ERROR);
    setError(value);
  }

  const setNotFound = (value) => {
    setStatus(NOT_FOUND);
    setError(value);
  }

  function renderResult(result) {
    if (!status) {
      return null;
    }
    if (status === SUCCESS) {
      return <Result value={result} key={result} />
    } else if (status === NOT_FOUND) {
      return <Alert variant="warning" className="text-center">{error}</Alert>
    } else {
      return <Alert variant="danger" className="text-center">{error}</Alert>
    }
  }

  return (
    <Container>
      <Navbar bg="dark" variant="dark">
        <Container>
          <Navbar.Brand href="/">Water Jug Riddle</Navbar.Brand>
        </Container>
      </Navbar>
      <Card>
        <Card.Body>
          <Card.Title>Fill the fields to get the problem solution</Card.Title>
          <Card.Body>
            <Container>
              <Card.Body>
                <RiddleForm onSuccess={setSuccess} onError={setStatusError} onNotFound={setNotFound} />
                {renderResult(result)}
              </Card.Body>
            </Container>
          </Card.Body>
        </Card.Body>
      </Card>
    </Container>
  );
}

export default App;
