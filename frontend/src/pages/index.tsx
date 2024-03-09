import { gql, useQuery } from "@apollo/client";

const TEST_QUERY = gql`
query TestQuery {
    test
}
`

const Index: React.FC = () => {
    const { loading, error, data } = useQuery(TEST_QUERY);

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

    return (
        <div>
            <h1>Your random number is: {data.test}</h1>
            <h1>Refresh the page to get a new one from the server!</h1>
        </div>
    )
}

export default Index;