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
import GaugeChart from 'react-gauge-chart';
import DonutChart from 'react-donut-chart';
import Chart from "react-apexcharts";
import Gauge from 'react-svg-gauge';
// import 'bootstrap/dist/css/bootstrap.css';
import Card from 'react-bootstrap/Card';

const StatisticsComponent = () => {
    const {userData} = useContext(UserContext);
    const [averageUtil, setAverageUtil] = useState(0.0);
    const [occupancy, setOccupancy] = useState({occupied: 0, total: 0});
    const [automated, setAutomated] = useState({automated: 0, manual: 0});
    const [yearlyData, setYearlyData] = useState([{}]);
    // const [yearlyData, setYearlyData] = useState([{date: 0, percentage: 0}]);
    
    const getG = () => {
        return 90;

    };

    const getColour = (val) => {
        if (val >= 0.80) {
            return '#00ff00';
        }
        else if (val >= 0.50) {
            return '#ffff00';
        }
        else {
            return '#ff0000';
        }
    };

    const formatYearlyData = (yd) => {
        var arr = new Array();

        yd.forEach(el => {
            arr.push({x: new Date(el.date).getTime(), y: el.percentage});
        });

        // return arr;
        
        return [{
                x: new Date("2022-09-01T00:00:00Z").getTime(),
                y: 76
              }, 
              {
                x: new Date("2022-08-01T00:00:00Z").getTime(),
                y: 56
              },
              {
                x: new Date("2022-07-01T00:00:00Z").getTime(),
                y: 35
              },
              {
                x: new Date("2022-06-01T00:00:00Z").getTime(),
                y: 72
              },
              {
                x: new Date("2022-05-01T00:00:00Z").getTime(),
                y: 60
              },
              {
                x: new Date("2022-04-01T00:00:00Z").getTime(),
                y: 100
              },
              {
                x: new Date("2022-03-01T00:00:00Z").getTime(),
                y: 76
              },
              {
                x: new Date("2022-02-01T00:00:00Z").getTime(),
                y: 67
              },
              {
                x: new Date("2022-01-01T00:00:00Z").getTime(),
                y: 84
              },
              {
                x: new Date("2021-12-01T00:00:00Z").getTime(),
                y: 14
              },
              {
                x: new Date("2021-11-01T00:00:00Z").getTime(),
                y: 72
              },
              {
                x: new Date("2021-10-01T00:00:00Z").getTime(),
                y: 95
              }];
    };

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
            // setAverageUtil(data.average.average.toFixed(1));
            setAverageUtil(data.average.average / 100.0);
            // setAverageUtil(0.8);
            // console.log(averageUtil);

            setOccupancy(data.occupancy);

            setAutomated(data.automated);

            // setYearlyData(data.yearly_utilisation);
            // "2022-10-01T00:00:00Z"
            // const ttt = new Date("2022-10-01T00:00:00Z").getTime();
            // console.log(ttt);
            setYearlyData(formatYearlyData(data.yearly_utilisation))
        });
    },[userData.token]);

    return (
        <div className={styles.statisticsContainer} class='overflow-auto'>
            <div className={styles.statisticsHeadingContainer}>
                <div className={styles.statisticsHeading}>Statistics</div>
            </div>
            <div style={{overflowY: 'scroll', height: '80vh'}}>
                <div class="row mx-2">
                    <div class="col-auto mr-2 mt-4 ">
                        <Card  style={{width: '750px', height: 'auto', textAlign: 'center'}} >
                            <Card.Body>
                                <Card.Title style={{fontSize: '1.5rem', fontWeight: 'bold'}}>
                                    Average Monthly Utilisation
                                </Card.Title>
                                <Chart 
                                    series={[{data: yearlyData}]}
                                    options= {{
                                        chart: {
                                            height: 350,
                                            type: 'line',
                                            zoom: {
                                                enabled: false
                                            }
                                        },
                                        dataLabels: {
                                            enabled: false
                                        },
                                        stroke: {
                                            curve: 'straight'
                                        },
                                        title: {
                                            text: undefined
                                        },
                                        grid: {
                                            row: {
                                                colors: ['#f3f3f3', 'transparent'], // takes an array which will be repeated on columns
                                                opacity: 0.5
                                            },
                                        },
                                        xaxis: {
                                            type: 'datetime',
                                        },
                                        yaxis: {
                                            min: 0,
                                            max: 100,
                                        },
                                        theme: {
                                            mode: 'light', 
                                            palette: 'palette10', 
                                            monochrome: {
                                                enabled: false,
                                                color: '#255aee',
                                                shadeTo: 'light',
                                                shadeIntensity: 0.65
                                            },
                                        }
                                    }}                  
                                    type="line" 
                                    height={350} 
                                />
                            </Card.Body>
                        </Card>
                    </div>

                    <div class="col-auto mr-2 mt-4">
                        {/* <Card style={{width: '18rem'}} > */}
                        <Card style={{width: '350px', textAlign: 'center'}} >
                            <Card.Body>
                                <Card.Title style={{fontSize: '1.5rem', fontWeight: 'bold'}}>
                                    Average Resource Utilisation
                                </Card.Title>
                                <GaugeChart id='GaugeChart2' 
                                    class="" 
                                    percent={averageUtil}
                                    nrOfLevels={2} 
                                    arcsLength={[averageUtil, (1-averageUtil)]} 
                                    colors={[getColour(averageUtil), '#edebeb']}
                                    arcPadding={0}
                                    arcWidth={0.3}
                                    cornerRadius={0}
                                    textColor={'#000000'}
                                    needleColor={"transparent"}
                                    needleBaseColor={"transparent"}
                                    fontWeight="900"
                                    fontSize="2.5rem"
                                    lineHeight='10px'
                                />
                            </Card.Body>
                        </Card>
                    </div>

                    <div class="col-auto mr-2 mt-4">
                        {/* <Card style={{width: '18rem'}} > */}
                        <Card style={{width: '500px', height: '400px', textAlign: 'center'}} >
                        {/* <Card> */}
                            <Card.Body>
                                <Card.Title style={{fontSize: '1.5rem', fontWeight: 'bold'}}>
                                    Current Office Occupancy
                                </Card.Title>
                                <DonutChart id='DonutChart' 
                                    data={[
                                        {
                                        label: 'Occupied',
                                        value: occupancy.occupied,
                                        },
                                        {
                                        label: 'Empty',
                                        value: (occupancy.total-occupancy.occupied),
                                        },
                                    ]}
                                    interactive={false}
                                    colors={['#00ff00', '#edebeb']}
                                    strokeColor={'transparent'}
                                    width={'500'}
                                />
                            </Card.Body>
                        </Card>
                    </div>

                    <div class="col-auto mr-2 mt-4">
                        {/* <Card style={{width: '18rem'}} > */}
                        <Card style={{width: '520px', height: '400px', textAlign: 'center'}} >
                        {/* <Card> */}
                            <Card.Body>
                                <Card.Title style={{fontSize: '1.5rem', fontWeight: 'bold'}}>
                                    Manual to Automated Bookings
                                </Card.Title>
                                <DonutChart id='DonutChart' 
                                    data={[
                                        {
                                        label: 'Automated',
                                        value: automated.automated,
                                        },
                                        {
                                        label: 'Manual',
                                        value: automated.manual,
                                        },
                                    ]}
                                    interactive={false}
                                    colors={['#00c3e3', '#ff4554']}
                                    strokeColor={'transparent'}
                                    width={'500'}
                                />
                            </Card.Body>
                        </Card>
                    </div>
                                    
                    

                    <div class="col-auto mr-2 mt-4">
                        {/* <Card style={{width: '18rem'}} > */}
                        <Card style={{width: '350px', textAlign: 'center'}} >
                            <Card.Body>
                                <Card.Title style={{fontSize: '1.5rem', fontWeight: 'bold'}}>
                                    Graph 3
                                </Card.Title>
                                <GaugeChart id='GaugeChart3' 
                                    class="" 
                                    nrOfLevels={2} 
                                    arcsLength={[0.3, 0.7]} 
                                    colors={[getColour(averageUtil), '#edebeb']}
                                    arcPadding={0}
                                    cornerRadius={0}
                                    textColor={'#000000'}
                                    needleColor={"transparent"}
                                    needleBaseColor={"transparent"}
                                />
                            </Card.Body>
                            
                        </Card>
                    </div>

                    

                </div>
                <div class='row mx-2'>
                    
                </div>
            </div>
            
        </div>
    )
}

export default StatisticsComponent;
