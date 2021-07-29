import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';
import { useState } from 'react';

function RiddleForm(props) {

  const [value, setValue] = useState({ x: 0, y: 0, z: 0 });

  const handleInputChange = e => {
    setValue({ ...value, [e.target.name]: parseInt(e.target.value) })
  }

  const onSubmit = (e) => {
    e.preventDefault();
    (async () => {
      const rawResponse = await fetch('/solve', {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(value)
      });
      if (rawResponse.status === 200) {
        const content = await rawResponse.json();
        props.onSuccess(content);
      } else if (rawResponse.status === 404) {
        const content = await rawResponse.text();
        props.onNotFound(content);
      } else {
        const content = await rawResponse.text();
        props.onError(content);
      }
    })();
  }

  const { x, y, z } = value;

  return (
    <Form onSubmit={onSubmit}>
      <Row>
        <Form.Group as={Col}>
          <Form.Label>X Gallon Jug</Form.Label>
          <Form.Control type="number" value={x} name="x" onChange={handleInputChange} placeholder="Please enter a value" />
        </Form.Group>
        <Form.Group as={Col}>
          <Form.Label>Y Gallon Jug</Form.Label>
          <Form.Control type="number" value={y} name="y" onChange={handleInputChange} placeholder="Please enter a value" />
        </Form.Group>
        <Form.Group as={Col}>
          <Form.Label>Desired Measure (Z)</Form.Label>
          <Form.Control type="number" value={z} name="z" onChange={handleInputChange} placeholder="Please enter a value" />
        </Form.Group>
      </Row>
      <Row>
        <Col>
          <Card.Body className="d-flex justify-content-center">
            <Button variant="primary" type="submit">
              Submit
            </Button>
          </Card.Body>
        </Col>
      </Row>
    </Form>
  );
}

export default RiddleForm;