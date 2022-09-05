import React, { useCallback } from 'react';
import Particles from "react-tsparticles";
import { loadFull } from 'tsparticles';

const Background = () =>
{
    const particlesInit = useCallback(async (engine) =>
    {
        await loadFull(engine);
    }, []);

    return (
        <Particles
            init={particlesInit}
            options = 
            {{
                fullScreen: {
                    enable: true,
                    zIndex: 1
                },
                particles: {
                    number: {
                        value: 10,
                        limit: 10,
                        density: {
                            enable: true,
                            value_area: 800
                        }
                    },
                    color: {
                        value: "#c3a3fe"
                    },
                    shape: {
                        type: "square",
                        stroke: {
                            width: 0,
                            color: "#000000"
                        },
                        polygon: {
                            nb_sides: 5
                        },
                    },
                    opacity: {
                        value: 1,
                        random: true,
                    },
                    size: {
                        value: 200,
                        random: true,
                    },
                    links: {
                        enable: false
                    },
                    move: {
                        enable: true,
                        speed: 1,
                        direction: "none",
                        random: false,
                        straight: false,
                        out_mode: "out",
                        bounce: false,
                        attract: {
                            enable: false,
                            rotateX: 600,
                            rotateY: 1200
                        }
                    }
                },
                detectRetina: true,
                fpsLimit: 60,
                background: {
                    color: "#b6c1fe"
                }
            }}
        />
    )
}

export default Background