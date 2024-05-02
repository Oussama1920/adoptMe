import './PetListStyle.css';
import Pet from './Pet' 


function PetList(props) {
    
    return <ul className="pets">
    <Pet type="bird" name="zakzouk" age="2ans" available="true" photo="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQET4cjElPT1W4e6ZB-1aQVT6UX2-l8VSCW_CzaMr9VQA&s" />
    <Pet type="dog" name="kahla" age="2ans" photo="https://cdn.britannica.com/70/234870-050-D4D024BB/Orange-colored-cat-yawns-displaying-teeth.jpg" />
    <Pet type="cat" name="katkout" age="2ans" photo="https://i.natgeofe.com/n/548467d8-c5f1-4551-9f58-6817a8d2c45e/NationalGeographic_2572187_square.jpg" />
    <Pet type="cat" name="katkout" age="2ans" photo="https://i.natgeofe.com/n/548467d8-c5f1-4551-9f58-6817a8d2c45e/NationalGeographic_2572187_square.jpg" />
    <Pet type="cat" name="katkout" age="2ans" photo="https://i.natgeofe.com/n/548467d8-c5f1-4551-9f58-6817a8d2c45e/NationalGeographic_2572187_square.jpg" />
    <Pet type="cat" name="katkout" age="2ans" photo="https://upload.wikimedia.org/wikipedia/commons/thumb/4/4d/Cat_November_2010-1a.jpg/640px-Cat_November_2010-1a.jpg" />
    <Pet type="cat" name="katkout" age="2ans" photo="https://upload.wikimedia.org/wikipedia/commons/thumb/4/4d/Cat_November_2010-1a.jpg/640px-Cat_November_2010-1a.jpg" />
    <Pet type="cat" name="katkout" age="2ans" photo="https://www.wfla.com/wp-content/uploads/sites/71/2023/05/GettyImages-1389862392.jpg?w=2560&h=1440&crop=1" />

    </ul>
}

export default PetList;