import { useState } from "react"


const FormField = ({type, label, fieldId, placeholder, required, children, validator, onStateChanged}) =>
{
    const [stateValues, SetStateValues] = useState({value : '', dirty : false, errors : []});

    HasChanged = e =>
    {
        e.preventDefault();

        const { label, required = false, validator = f => f, onStateChanged = f => f } = this.props;

        const value = e.target.value;
        const isEmpty = value.length === 0;
        const requiredMissing = stateValues.dirty && required && isEmpty;

        let errors = [];

        if(requiredMissing)
        {
            errors = [...errors, `${label} is required`];
        }
        else if('function' === typeof validator)
        {
            try
            {
                validate(value);
            }
            catch(e)
            {
                errors = [...errors, e.message];
            }
        }

        SetStateValues(({dirty = false}) => ({value, errors, dirty: !dirty || dirty}), () => onStateChanged(stateValues));
    }
}

export default FormField;