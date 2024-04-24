import { Paper } from "@mui/material";
import ClassList from "../../components/ClassList";
import { gql, useQuery } from "@apollo/client";


const CLASSES_QUERY = gql`
query ClassesQuery {
    classes {
        classInfo {
            university_name
            name
            description
        }
    }
}
`

const Classes: React.FC = () => {
    const { loading, error, data } = useQuery(CLASSES_QUERY);

    if (loading) {
        return (
            <h1>Loading</h1>
        )
    }

    if (error) {
        return (
            <h1>{error.message}</h1>
        )
    }

    const outData = []
    for (const part of data.classes) {
        outData.push(part.classInfo);
    }

    return (
        <Paper elevation={3}>
            <ClassList classes={outData} />
        </Paper>
    )
}

export default Classes;