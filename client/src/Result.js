import { useEffect, useState } from 'react';
import Jug from './Jug';
import Button from 'react-bootstrap/Button';
import ButtonGroup from 'react-bootstrap/ButtonGroup';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import ProgressBar from 'react-bootstrap/ProgressBar';
import Card from 'react-bootstrap/Card';

function Result({ value }) {

    const [step, setStep] = useState(0);

    useEffect(() => {
        setStep(0);
    }, [value])

    if (!value.jugs || !value.steps || step > value.steps.length - 1) {
        return null;
    }

    const { left, right } = value.jugs;
    const total = value.steps.length - 1;
    const currentStep = value.steps[step];

    return (
        <Container>
            <Row>
                <Col>
                    <Jug capacity={left.capacity} value={currentStep.left} title={left.name} />
                </Col>
                <Col className="d-flex align-items-center">
                    <h4 className="d-flex justify-content-center" style={{ width: '100%' }}>{currentStep.name}</h4>
                </Col>
                <Col>
                    <Jug capacity={right.capacity} value={currentStep.right} title={right.name} />
                </Col>
            </Row>
            <Row>
                <Col>
                    <h2 className="d-flex justify-content-center">{step} / {total}</h2>
                </Col>
            </Row>
            <Row>
                <Col>
                    <Card.Body>
                        <ProgressBar now={step * 100 / total} />
                    </Card.Body>
                </Col>
            </Row>
            <Row>
                <Col className="d-flex justify-content-center">
                    <ButtonGroup>
                        <Button disabled={step === 0} onClick={() => setStep(step - 1)}>Prev</Button>
                        <Button disabled={step === total} onClick={() => setStep(step + 1)}>Next</Button>
                    </ButtonGroup>
                </Col>
            </Row>
        </Container>
    )
}

export default Result;