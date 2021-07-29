import { useSpring, animated } from 'react-spring';
import Card from 'react-bootstrap/Card';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

function Jug({ capacity, value, title }) {

    const filled = value * 100 / capacity;
    const styles = useSpring({ height: `${filled}%`, width: '100%', backgroundColor: '#0d6efd', borderRadius: 5 })

    return (
        <Container>
            <Card.Body>
                <Row>
                    <Col className="d-flex justify-content-center">
                        <h3>{title}</h3>
                    </Col>
                </Row>
                <Row>
                    <Col className="d-flex justify-content-center">
                        <Card style={{ height: 25, width: 25, border: 'solid', borderRadius: 5 }}>
                        </Card>
                    </Col>
                </Row>
                <Row>
                    <Col className="d-flex justify-content-center">
                        <Card style={{ height: 150, width: 150, border: 'solid', borderRadius: 10, display: 'flex', justifyContent:'flex-end' }}>
                            <animated.div style={styles}>
                            </animated.div>
                        </Card>
                    </Col>
                </Row>
                <Row>
                    <Col className="d-flex justify-content-center">
                        <h4>{value} / {capacity}</h4>
                    </Col>
                </Row>
            </Card.Body>
        </Container>
    );
}

export default Jug;