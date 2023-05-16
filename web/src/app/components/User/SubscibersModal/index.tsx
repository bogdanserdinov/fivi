import { Avatar } from "@components/common/Avatar"
import { Modal } from "@components/common/Modal"

import './index.scss'

export const UserSubscribersModal: React.FC<{ subscribers: any, setIsOpenModal: React.Dispatch<React.SetStateAction<boolean>> }> = ({ subscribers, setIsOpenModal }) => {
    return (
        <Modal setIsOpenModal={setIsOpenModal}>
            <div className="subscribers">
                {subscribers ?
                    subscribers.map((subscriber: any) =>
                        <div className="subscribers__item">
                            <div className="subscribers__item__info">
                                <Avatar size={50} photo={subscriber.avatar} />
                                <p className="subscribers__item__nickname">{subscriber.nickname}</p>
                                {!subscriber.isSubscribe &&
                                    <>
                                    &#8226;
                                    <button className="subscribers__item__subscribe">Підписатися</button>
                                    </>
                                }
                            </div>
                            <button className="subscribers__item__delete">
                                Видалити
                            </button>
                        </div>
                    ) :
                    <p>Ще нема підписніків</p>
                }
            </div>

        </Modal>
    )
}