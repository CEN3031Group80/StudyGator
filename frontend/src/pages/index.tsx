import { gql, useQuery } from "@apollo/client";

const TEST_QUERY = gql`
query TestQuery {
    me {
        id
        avatarURL
    }
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
            <h1>Your user id is: {data.me.id}</h1>
        </div>
    )
}

export default Index;