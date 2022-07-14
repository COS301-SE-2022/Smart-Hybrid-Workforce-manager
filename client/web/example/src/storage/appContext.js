//Code sourced from: // https://github.com/4GeeksAcademy

import React, { useState, useEffect } from "react";
import getState from "./data.js";

export const Context = React.createContext(null);

const globalContext = PassedComponent => {
	const storageWrapper = props => {
		//this will be passed as the contenxt value
		const [state, setState] = useState(
			getState({
				getAuthData: () => state.auth_data,
                setAuthData: updatedData => setState({auth_data: Object.assign(state.auth_data, updatedData)
                })
			})
		);

		// The initial value for the context is not null anymore, but the current state of this component,
		// the context will now have a getStore, getActions and setStore functions available, because they were declared
		// on the state of this component
		return (
			<Context.Provider value={state}>
				<PassedComponent {...props} />
			</Context.Provider>
		);
	};
	return storageWrapper;
};

export default globalContext;
