import { user } from "@/mocked/user"
import { Avatar } from "@components/common/Avatar"
import { useState } from "react"
import backButtonIcon from '@img/backButtonIcon.png'

import './index.scss'
import { Link } from "react-router-dom"

export const UserEdit = () => {
    const [photo,setPhoto]=useState<string>(user.avatar)
    const [nickName, setNickName] = useState(user.nickname)
    
    const handleFileChange = (e:any) => {
         if (e.target.files?.length) {
                setPhoto(URL.createObjectURL(e.target.files[0]));
        }
    }
     const handleNicknameChange = (e:any) => {
        setNickName(e.target.value)
    }
    
    return (
        <div className="user-edit">
            <Link to={`/user/${user.id}`} className="user-edit__back-button">
                <img src={backButtonIcon} alt="back button" className="user-edit__back-button__image"/>
            </Link>
            
            <div className="user-edit__profile">
                <label htmlFor="user-edit" className="user-edit__file-label">
                    <Avatar photo={photo} size={160} />
                    <p className="user-edit__file-label__text">Змінити фото</p>
                </label>
                <input
                    type="file"
                    accept="image/png, image/jpeg"
                    onChange={handleFileChange}
                    id="user-edit"
                    className="user-edit__file-input"
                />
            </div> 
            <div className="user-edit__field">
                <label className="user-edit__field__label">Змінити нік</label>
                <input className="user-edit__field__input" type="text" value={nickName} onChange={handleNicknameChange} />
            </div>
        </div>
    )
}