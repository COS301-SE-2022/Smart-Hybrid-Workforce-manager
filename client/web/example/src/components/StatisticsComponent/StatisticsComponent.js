import styles from './statistics.module.css';
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';
import { MdEdit, MdPersonAdd, MdClose } from 'react-icons/md';
import { useContext, useEffect, useRef, useState } from 'react';
import { IoIosArrowBack, IoIosArrowForward } from 'react-icons/io';
import { AiOutlineUsergroupAdd } from 'react-icons/ai';
import { FaSave } from 'react-icons/fa';
import { BsThreeDotsVertical } from 'react-icons/bs';
import { EditTeamForm } from '../Team/EditTeam';
import { AddTeamForm } from '../Team/AddTeam';
import { EditUserPanel } from '../User/EditUser';
import { UserContext } from '../../App';
import GaugeChart from 'react-gauge-chart'
import Gauge from 'react-svg-gauge'
import 'bootstrap/dist/css/bootstrap.css';
import Card from 'react-bootstrap/Card';

const StatisticsComponent = () => {
    const {userData} = useContext(UserContext);
    const [averageUtil, setAverageUtil] = useState(0.0);
    // var averageUtil;
    var test4 = 57;
    const getG = () => {
        return 90;

    };

    const getColour = (val) => {
        if (val >= 80) {
            return '#00ff00';
        }
        else if (val >= 50) {
            return '#ffff00';
        }
        else {
            return '#ff0000';
        }
    }

    const val = getG();

    useEffect(() =>
    {
        fetch("http://localhost:8080/api/statistics/all", 
        {
            method: "POST",
            mode: "cors",
            body: JSON.stringify({
            }),
            headers:{
                'Content-Type': 'application/json',
                'Authorization': `bearer ${userData.token}` 
            }
        }).then((res) => res.json()).then(data => 
        {
            console.log(data);
            // averageUtil = data.average.average;
            setAverageUtil(data.average.average.toFixed(1));
            // console.log(averageUtil);
        });
    },[userData.token]);

    return (
        <div className={styles.statisticsContainer}>
            <div className={styles.statisticsHeadingContainer}>
                <div className={styles.statisticsHeading}>Statistics</div>
            </div>
            <div class="row mx-2">
                <div class="col-auto mr-2 mt-4 ">
                    <Card  >
                        <Card.Body>
                            {/* <Card.Title>Graph 1</Card.Title> */}
                            <Gauge id='GaugeChart1' class="" nrOfLevels={20} value={averageUtil} valueFormatter={number => `${number}%`} label={'Resource Utilisation'} color={getColour(averageUtil)} />
                        </Card.Body>
                        
                    </Card>
                </div>

                <div class="col-auto mr-2 mt-4">
                    <Card /*style={{width: '18rem'}}*/ >
                        <Card.Body>
                            <Card.Title>Graph 2</Card.Title>
                            <GaugeChart id='GaugeChart2' class=""/>
                        </Card.Body>
                        
                    </Card>
                </div>
            </div>
            
        </div>
    )
}

export default StatisticsComponent;
