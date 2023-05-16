import { Avatar } from "@components/common/Avatar"
import { Modal } from "@components/common/Modal"
import { Link } from "react-router-dom"

import './index.scss'

export const UserSubscribesModal: React.FC<{
    subscribes: any,
    setIsOpenModal: React.Dispatch<React.SetStateAction<boolean>>
}> = ({ subscribes, setIsOpenModal }) => {
    const rejectSubscribe = () => {
        
    }

    return (
       <Modal setIsOpenModal={setIsOpenModal}>
            <div className="subscribes">
                {subscribes ?
                    subscribes.map((subscribe: any) =>
                        <div className="subscribes__item">
                            <Link className="subscribes__item__info" to={subscribe.id}>
                                <Avatar size={50} photo={subscribe.avatar} />
                                <p className="subscribes__item__nickname">{subscribe.nickname}</p>
                            </Link>
                            <button className="subscribes__item__reject-subscribes">
                               Підписки
                            </button>
                        </div>
                    ) :
                    <p>Ще нема підписок</p>
                }
            </div>
        </Modal>
    )
}