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
import 'bootstrap/dist/css/bootstrap.css';
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
        if (val >= 80) {
            return '#00ff00';
        }
        else if (val >= 50) {
            return '#ffff00';
        }
        else {
            return '#ff0000';
        }
    };

    const formatYearlyData = (yd) => {
        var arr = new Array();

        yd.forEach(el => {
            arr.push({x: new Date(el.date).getTime(), y: el.percentage.toFixed(0)});
        });

        return arr;
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
            setAverageUtil(data.average.average.toFixed(1));
            setOccupancy(data.occupancy);
            setAutomated(data.automated);
            setYearlyData(formatYearlyData(data.yearly_utilisation));
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
                        <Card  style={{width: '1245px', height: 'auto', textAlign: 'center'}} >
                            <Card.Body>
                                <Card.Title style={{fontSize: '1.5rem', fontWeight: 'bold'}}>
                                    Average Monthly Utilisation
                                </Card.Title>
                                <Chart 
                                    series={[{name: 'utilisation percentage', data: yearlyData}]}
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
                                            palette: 'palette3',
                                        }
                                    }}                  
                                    type="line" 
                                    height={350} 
                                />
                            </Card.Body>
                        </Card>
                    </div>

                    <div class="col-auto mr-2 mt-4">
                        <Card style={{textAlign: 'center', width: '400px', paddingBottom: '-50px'}} >
                            <Card.Body style={{paddingBottom: '0px'}}>
                                <Card.Title style={{fontSize: '1.5rem', fontWeight: 'bold'}}>
                                    Average Resource Utilisation
                                </Card.Title>
                                <Chart 
                                    series={[averageUtil]}
                                    options= {{
                                        chart: {
                                            type: 'radialBar',
                                            offsetY: -40,
                                        },
                                        labels: ['Occupied', 'Empty'],
                                        dataLabels: {
                                            enabled: false
                                        },
                                        plotOptions: {
                                            radialBar: {
                                                startAngle: -120,
                                                endAngle: 120,
                                                track: {
                                                    background: '#f0f0f0',
                                                },
                                                dataLabels: {
                                                    name: {
                                                        show: false,
                                                    },
                                                    value: {
                                                        fontSize: '2rem',
                                                        offsetY: -5,
                                                    }
                                                }
                                            }
                                        },
                                        colors: [getColour(averageUtil)],
                                        grid: {
                                            padding: {
                                                top: -15,
                                                bottom: 0,
                                            }
                                        },
                                    }}     
                                    type='radialBar'    
                                    height='500px'  
                                    paddingBottom='-50px'
                                    // width='300px'
                                />
                            </Card.Body>
                        </Card>
                    </div>

                    <div class="col-auto mr-2 mt-4">
                        {/* <Card style={{width: '18rem'}} > */}
                        <Card style={{width: '400px', height: 'auto'}} >
                        {/* <Card> */}
                            <Card.Body>
                                <Card.Title style={{fontSize: '1.5rem', fontWeight: 'bold', textAlign: 'center'}}>
                                    Current Office Occupancy
                                </Card.Title>
                                <Chart 
                                    series={[occupancy.occupied, (occupancy.total-occupancy.occupied)]}
                                    options= {{
                                        chart: {
                                            type: 'donut',
                                        },
                                        labels: ['Occupied', 'Empty'],
                                        dataLabels: {
                                            enabled: false
                                        },
                                        plotOptions: {
                                            pie: {
                                                donut: {
                                                    labels: {
                                                        show: true,
                                                        total: {
                                                            showAlways: false,
                                                            show: true,
                                                        },
                                                    }
                                                }
                                            }
                                        },
                                        theme: {
                                            mode: 'light', 
                                            palette: 'palette3', 
                                        },
                                        legend: {
                                            show: true,
                                            offsetX: 25,
                                        },
                                        grid: {
                                            padding: {
                                                bottom: 7,
                                            }
                                        },
                                    }}     
                                    type='donut'      
                                    width='425'
                                />
                            </Card.Body>
                        </Card>
                    </div>

                    <div class="col-auto mr-2 mt-4">
                        {/* <Card style={{width: '18rem'}} > */}
                        <Card style={{width: '400px', height: 'auto', marginBottom: 25}} >
                        {/* <Card> */}
                            <Card.Body>
                                <Card.Title style={{fontSize: '1.5rem', fontWeight: 'bold', textAlign: 'center'}}>
                                    Bookings by Type
                                </Card.Title>
                                <Chart 
                                    series={[automated.automated, automated.manual]}
                                    options= {{
                                        chart: {
                                            type: 'donut',
                                        },
                                        labels: ['Automated', 'Manual'],
                                        dataLabels: {
                                            enabled: false
                                        },
                                        plotOptions: {
                                            pie: {
                                                donut: {
                                                    labels: {
                                                        show: true,
                                                        total: {
                                                            showAlways: false,
                                                            show: true,
                                                        }
                                                    }
                                                }
                                            }
                                        },
                                        theme: {
                                            mode: 'light', 
                                            palette: 'palette7',
                                        },
                                        legend: {
                                            show: true,
                                            offsetX: 25,
                                        },
                                        grid: {
                                            padding: {
                                                bottom: 14,
                                            }
                                        },
                                    }}     
                                    type='donut'          
                                    width='425'  
                                />
                            </Card.Body>
                        </Card>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default StatisticsComponent;
