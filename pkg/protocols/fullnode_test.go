package protocols_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chia-network/go-chia-libs/pkg/protocols"
	"github.com/chia-network/go-chia-libs/pkg/streamable"
)

func TestRespondPeers(t *testing.T) {
	// Has one peer in the list
	// IP 1.2.3.4
	// Port 8444
	// Timestamp 1643913969
	hexStr := "0000000100000007312e322e332e3420fc0000000061fc22f1"

	// Hex to bytes
	encodedBytes, err := hex.DecodeString(hexStr)
	assert.NoError(t, err)

	rp := &protocols.RespondPeers{}

	err = streamable.Unmarshal(encodedBytes, rp)
	assert.NoError(t, err)

	assert.Len(t, rp.PeerList, 1)

	pl1 := rp.PeerList[0]
	assert.Equal(t, "1.2.3.4", pl1.Host)
	assert.Equal(t, uint16(8444), pl1.Port)
	assert.Equal(t, uint64(1643913969), pl1.Timestamp)

	// Test going the other direction
	reencodedBytes, err := streamable.Marshal(rp)
	assert.NoError(t, err)
	assert.Equal(t, encodedBytes, reencodedBytes)
}

func TestRequestBlock(t *testing.T) {
	blockHeight := uint32(2347)
	hexStr := "00000000000000000000000000000000000040340000092b000000000000000000000001adcac4c82bc188b377ded9d77e4ba90ec871f9c0b8e4798a1d23a65900e0c10d3d9bd31d291abb6b4c6ed7c151032ddf690745b22d2f6f10f2c0ac3f8e14f5150e18f115f00182fd324dffdf8c57792e9d26adc9841ea0d8214afae9f44c5bd3f06af06bcaee17dca035be794c36c423288cef2fdf5900b66a2f7e9c163da148e9101ca260b084f678b4ae0ca5772b5750b7f44394efcfc351800f21eac65d900ddebbbf8e47732000000100b18870d5b5fcd7baa9445c1f8458beac4dcd10fa9254646a5c6b11d018be82a803487d8518448052600935b18943f7e7db758db3ac1631018e3d18336505e452c768506d44094bf209c2b015ff2cedfd7aa6e9256e13b4652adb8c1452c501a933404761806f4e11ce2e138602fb17ee0383c071b8fb7f47679501b77a1d53c18093ebf53269a43650bb78737e3da4150d34927b5c3cfd994d77b34f061376d104e64dfafbf9eb313f0b0aa88059b7d95af5c8ee0e4ea4737a7172109fa990e5587917e5c6b07237de592d4309c319095311b920cf5963c4e7243c9de9fc3d3cadc69bfcc503fe28551e6886f4a308cc9383eb552516908323ca4bb7a006c70d01c188b377ded9d77e4ba90ec871f9c0b8e4798a1d23a65900e0c10d3d9bd31d2900000000056000000200ae256619f34f64c60b9c3821ea69e178cc601e49cc4b246176c17c51ed694b8584da0393e82142029f717214e0b44a31aaaf9875f27079f40de19ddec6b3360269a5161c2f2769f675f8e7ce8b457fa7bcc0a86e1f9330b9dd4e559b38137a0d010083b3a189e345f2af24515e39ea92b13bacd6dd8c14b257725bd93cdbb63ea88f135e3acf7e6ae5b960a6cd8db0cb3d020512990bba2ce011ea58cad1262b28e0db72e4309c7a2695908bfa3aefeb89335407626f7da8dbe10ca6f926de15a981c188b377ded9d77e4ba90ec871f9c0b8e4798a1d23a65900e0c10d3d9bd31d290000000005cac4c80100424506b6b6957111f455a922fcc02a6456170c2d2a9a77a58f35d95bc14e0d91be5efb9d73b4c3142b2d9ca8445697164088fb2f3b7569239820e62166c58b2009ac1c20347fec68b93fbeba55080cdf09b671b17b29360464a22785fff5b81a010001103b14520e67df9edf58dcd3028c37d779b20e2a9a1d8870cd8d34b81305e92d0000000000028f1202001d872926b6b4512e9f86240409a4dcb529a789f33bd45d5c3403d8557d20731122d457feb0a4c2ba79411e4213a14a9437a4fa71c88bdf2f60085c0a6b154f2790325aeade83148ce0138a1df6479010ad831be24942d2effeca810d0f1baf37020195489a1064ce6366a92284e10eb4bacd284f7d7e085ecd62782fc7b1703d4bd395f65a469c5c37ae8a16b0756a1f5ad80407bf1c588749e90f68150ef14a407551eecccc329849f48e161aebf55d7085129e307588d1853e5e037fab7259df29cf4d9ef9dac40874f3a589f8eeaf10825bcb9e9b90074adc72cf6f3382ae65e900000000001b78cf00003e7378f866dabdd35ee5e2c5cfb7115b725c2e90df25dbba2590ff6d4b4e56feebdc774d01aee49e95e529e6979e6f47259c3611b3396fbdca3a2c4dd82fac68836eee4a71bc7f75a12f3febcdc1f80bf201b2ff5b9d994df9c3fda9df0a5b17010001adaa67a9a870fd370fe16fcbb65c4be294dfb6579600323d6e1e597b11256dee00000000001b78cf03005e37c168d5761771d986f4766c6972ef9b0653f3adf917d335e7cefd790cbecddecf021f9df463a6a801e8eedc73aa569f5c027bdec416813f2ff334c88fe72051f1568e27811b011d35da29099004eb09ba6a038844832bd725009d18a1b81701000001000000006403002c6a887316e81f4273ab2beb3ac84f31825a31f4afd0e723545b75acea10caf64b5dda4ddffa13b985086b2ee88fa5a1d4748d34f9b8ccd8b40418a8291d3e0491b4ad5433d7e9e1902f9e73398cff543372d6ba60bf07c21161e236da73f7010e05010000000064000057b0969c023691431262bd495c50cb005f3388a4750566a98b35761827e5f94a5673cb0bd89013313db595e51dc0c19cd3b426806d79420320f4e84031bd3f0befba686e74483ddc8ea239d27a37048474a09dd920d0b5f2fb91571f63d7762401000101020000017e0000d09482e77fb1dcf7c0229a46692a4f38f991688149f676a876ba0c3e166f3994646381a763404f5584d76283ecafc2fe678a0a81b1e75e73a36d19159182600395ac30168044f5f427071e76d97dd4b70eab06be6fa3ac24793a4488780e6509010000000000000091b4ab6fab4f3e3f7912e28e385a195a4a4e65b325bfbe313114da885eaf74e642179d0200889b70e2b2032c4d44406a4e9177f4158fb7609b930876137437ef3de3af1c8e7e8500273ab77a3f6e40cceccd5d491ac291ec8af137c100aff9504007e1de0aeb572d73cb348d529da8defef95b8bd815e2f2814b7b528596fb1a6a48f26b130100000000000001b454bca04db73ba1c4c19c470866e511a71a73600837ac8467cba42b57705132b16c050100908b4e5520fdf9edaaf02fa81e54a1b59aacc90195ed7f96456b56f5d606b2d70d99f3270c8de66a19526d221f1b003e23e4b90def8b62f2c764861c61d724779fa7d397e5ab91c30106fe4957e674a3d2ff65c06d1808a91b32c23c08c9c312010000020000017e02001f906ed8f095bad9d70c3d5b792b7b3d6a50ee8f6f3fab32876e99910d4be748f01e450030e237dc7aef7380d266179a98a6305125e1f3ac2cacab93dde0402e6e4ae45d445d9cbe59026d1229cf241bd7e8486dcf3d20a2d2f01e48370ce14502010000000000061ae4907e5e332380a2689e77ea9a485fbb5d6a3dd84c5aa179d3d9a003fecc43f35b270300aca0158c98ff464f7942257f901ea60f2a95ad1f9376c17b42a0bba97bd86e5b65ed9e93526334c6ed0ceb1b2032d2af24547573569c2801fb2d8f42b0ad8c5daff72f0bc62dc2503f9385f00c719a7ddf9261166fef6114afbea3966ce1af050100000000000012504890a945b69bc719b5cc9c5059a348d5034c5401f0267d6a1128837178a10b3afce301002b418db30870754b5d0abaa3a4fd1060591a85ff964997ca4d29e42a71d019c22e1e025397642a89c0e13bb1f742c5ed97de55f97c0ab63e640e10be2b2ebe1bcbd629b8906cbb0420a1f5ab1583090ec08a341217d85756ccf2815846f7fd2102010001020000017e02009a50a6536d487c9a7fad852bc3ab9ed2dca8d633164a647f2cdf20468fb02ab4445dd6ce85d16641aee46e1340a530e0917caadc0f5ec9fc7fc1b928698c4a12ff306deb3689d18bf6c90af772bba1151cb1b82c4f319b86b0def633bd80c51805000000000000061ae4d251d5e5e51661e483dd39862ecfbed6ef3a48529b2b522a431c67862d9bb6548b0200d42a2de42e351a2b5b0c3e9aa3adc3806033f233147e1cc8e77984601e2fdb888712f45bdec670e65b9a38825f0e84a91ae4606b548a55a19cb4336dd2e5b569ff6867f1c28714cfb06314849e55896a025ae0d57c751864b5c8ca1d35e3245701000000000000125048e8c5405ea099da345f6f8fe9d092f9ea51b6a7fe48b37be3891b05c051db9393e101006a30c805895b64dc89886fe508c3c3f0fa2a25e4bc17f9be12a37ec32e03ce7c5548130ff2bc98e8a19452e5da259ba95b1c3b7e687463dc70e704948b9709097331606e05dbd94aea96017668943475f160cc3e91de3448dde4ed81cec2d0130100008416fadbf6c13544d2be22bf69ac81679c237e3852972af2694c1c26466025ce2a0b38d5feda302fb27579e85bbe01ae70ea0871d20601d0ed7f3abc25ebf9989825bc9b4470bf03fe10007531435dc7b5332552800aa120b36f3baa02abf9e1cea41d2961b6132fcaea28ca35b425cb9bb2a23c71fa8af3c44509dfbb4e64670000000001aee44571244f660c4c22e0e2f85d51c2bbdd75e8217e04afa159b603739aa0ab133cae0ef53fbcb5a2ad59cd976bffdf18cd2cd96c64790c9a11549a873e99db22e6ce56e2105ea2c37a6d1cafc1cfc3cc03bdcc8eea62bee18ed99de59fb828cea41d2961b6132fcaea28ca35b425cb9bb2a23c71fa8af3c44509dfbb4e64670000000000000000000000000000000000000000000000000000000003a2c7c9b14fca84c30a94fddee8dbd01327103f6f3db7b45ed493603ff39edd2c0702c1404ebdbf1216619572fe54d4320919a60be60dea18928fbc108e638a7edf8df714ce1b708dada423602913cd59be9e018dc128ed51c32e959e87a2a3e603cb49000000000000000000"

	// Hex to bytes
	encodedBytes, err := hex.DecodeString(hexStr)
	assert.NoError(t, err)

	rb := &protocols.RespondBlock{}

	err = streamable.Unmarshal(encodedBytes, rb)
	assert.NoError(t, err)

	assert.Equal(t, rb.Block.RewardChainBlock.Height, blockHeight)

	// Test going the other direction
	reencodedBytes, err := streamable.Marshal(rb)
	assert.NoError(t, err)
	assert.Equal(t, encodedBytes, reencodedBytes)
}
