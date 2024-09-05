import MapComponent from "../components/MapComponents/MapComponent.jsx";


const Map = () => {
    const locations = [
        { lat: 37.7749, lng: -122.4194 }, // San Francisco
        { lat: 40.7128, lng: -74.0060 },  // New York
        { lat: 51.5074, lng: -0.1278 },   // London
        { lat: 22.518, lng: 88.3832 },    // Paris
    ];
    return (
        <div>
            <h1>Geolocation Map</h1>
            <MapComponent locations={locations} />
        </div>
    );
}
export default Map;