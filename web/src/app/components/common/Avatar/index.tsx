export const Avatar: React.FC<{ size: number; photo: string }> = ({ size, photo }) =>
    <img
        src={photo}
        alt="avatar"
        style={{ width: `${size}px`, height: `${size}`, borderRadius: '50%' }}
    />;
